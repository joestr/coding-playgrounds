module main

import vweb

const (
	port = 8082
)

pub struct App {
pub mut:
	vweb vweb.Context // TODO embed
	cnt int
}

fn main() {
	println('noice')
	vweb.run<App>(port)
}

pub fn (app mut App) init() {
	app.vweb.handle_static('.')
}

pub fn (app mut App) json_endpoint() {
	app.vweb.json('{"a": 3}')
}

pub fn (app mut App) json2() {
	app.vweb.json('{"array": ["complex"], "xml": "none", "data": {"complex": true}}')
}

pub fn (app mut App) index() {
	app.cnt++
	$vweb.html()
}

pub fn (app mut App) reset() {
}

pub fn (app mut App) text() {
	app.vweb.text('Hello world')
}

pub fn (app mut App) cookie() {
	app.vweb.text('Headers:')
	app.vweb.set_cookie('cookie', 'test')
	app.vweb.text(app.vweb.headers)
	app.vweb.text('Text: hello world')
}