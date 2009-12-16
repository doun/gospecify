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

import "container/list";

type example interface {
	Run(Reporter, func(Example));
}

type exampleCollection struct {
	examples *list.List;
}

func makeExampleCollection() *exampleCollection { return &exampleCollection{list.New()}; }

func (self *exampleCollection) Add(e example) {
	self.examples.PushBack(e);
}

func (self *exampleCollection) Init(runner *runner) {
	for i := range self.examples.Iter() {
		if ex, ok := i.(*complexExample); ok {
			runner.currentExample = ex;
			ex.Init();
		}
	}
}

func (self *exampleCollection) Run(reporter Reporter, before func(Example)) {
	for i := range self.examples.Iter() {
		if ex, ok := i.(example); ok {
			ex.Run(reporter, before);
		}
	}
}