
// LoL Cruncher - A Historical League of Legends Statistics Tracker
// Copyright (C) 2015  Jason Chu (1lann) 1lanncontact@gmail.com

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package controllers

import (
	"github.com/revel/revel"
	"cruncher/app/models/query"
)

type View struct {
	*revel.Controller
}

func (c View) Index() revel.Result {
	return c.Render()
}

func (c View) Test() revel.Result {
	c.RenderArgs["errorCode"] = "ErrNotFound"
	return c.RenderTemplate("errors/database.html");
}

func (c View) Request() revel.Result {
	// func GetStats(name string, region string) (string, dataFormat.Player, error) {
	name, player, err := query.GetStats("1lann", "na")
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderTemplate("errors/database.html");
	}
	c.RenderArgs["player"] = player
	c.RenderArgs["name"] = name
	return c.Render()
}
