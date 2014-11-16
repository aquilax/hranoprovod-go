package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func readChannels(parser *Parser) (*NodeList, error) {
	nodeList := NewNodeList()
	for {
		select {
		case node := <-parser.nodes:
			nodeList.push(node)
		case breakingError := <-parser.errors:
			return nil, breakingError
		case <-parser.done:
			return nodeList, nil
		}
	}
}

func TestParser(t *testing.T) {
	Convey("Given new parser", t, func() {
		parser := NewParser(NewDefaultParserOptions())
		Convey("It completes successfully on empty string", func() {
			go parser.parseStream(strings.NewReader(""))
			nodeList, error := readChannels(parser)
			So(len(*nodeList), ShouldEqual, 0)
			So(error, ShouldBeNil)
		})

		Convey("It processes valid node", func() {
			file := `2011/07/17:
  el1: 1.22
  ел 2:  4
  el/3:  3`
			go parser.parseStream(strings.NewReader(file))
			nodeList, err := readChannels(parser)
			So(len(*nodeList), ShouldEqual, 1)
			So(err, ShouldBeNil)
			node := (*nodeList)["2011/07/17"]
			So(node.header, ShouldEqual, "2011/07/17")
			elements := node.elements
			So(elements, ShouldNotBeNil)
			So(len(*elements), ShouldEqual, 3)
			So((*elements)[0].name, ShouldEqual, "el1")
			So((*elements)[0].val, ShouldEqual, 1.22)
			So((*elements)[1].name, ShouldEqual, "ел 2")
			So((*elements)[1].val, ShouldEqual, 4.0)
			So((*elements)[2].name, ShouldEqual, "el/3")
			So((*elements)[2].val, ShouldEqual, 3.0)
		})
	})
}
