// Copyright 2013 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routers

// AboutRouter serves about page.
type AboutRouter struct {
	baseRouter
}

// Get implemented Get method for AboutRouter.
func (this *AboutRouter) Get() {
	this.Data["IsAbout"] = true

	// Get language.
	curLang, _ := this.Data["LangVer"].(langType)
	this.TplNames = "about_" + curLang.Lang + ".html"
}
