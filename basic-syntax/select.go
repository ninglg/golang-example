package main

func main() {
	/*
	select可同时监听多个channel的读/写
	执行select时，若只有一个case通过，则执行该case
	若有多个，则随机执行一个case
	若所有都不满足，则执行default，若无default，则等待
	可使用break跳出select
	 */
}