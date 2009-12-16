/*
Copyright (c) 2009 Samuel Tesla <samuel.tesla@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package specify

type runner struct {
	examples *exampleCollection;
	currentExample *complexExample;
}

func makeRunner() *runner { return &runner{examples:makeExampleCollection()}; }

func (self *runner) Before(block func(Example)) {
	if self.currentExample == nil { return; }
	self.currentExample.AddBefore(block);
}

func (self *runner) Describe(name string, block func()) {
	self.examples.Add(makeComplexExample(name, block));
}

func (self *runner) It(name string, block func(Example)) {
	if self.currentExample == nil { return; }
	self.currentExample.Add(makeSimpleExample(name, block));
}

func (self *runner) Run(reporter Reporter) {
	self.examples.Init(self);
	self.examples.Run(reporter, func(Example){});
	reporter.Finish();
}
