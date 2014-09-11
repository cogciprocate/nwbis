enum Direction {
    North,
    East,
    South,
    West
}

enum Color {
  Red = 0xff0000,
  Green = 0x00ff00,
  Blue = 0x0000ff
}

enum Shape {
    Circle(Point, f64),
    Rectangle(Point, Point)
}

fn main() {

	let mut x: int;
	x = 4;

	if x == 5i {
		println!("x is five!");
	} else {
    	println!("x is not five :(");
	}

	let a = if x == 5i { 10i } else { 15i };
	
	let (y, z) = (1i, 2i);
    println!("hello?");
    println!("x = {}; y = {}; z = {}; a = {}", x, y, z, a);
    print_number(a, add_one(x));
    tuples();
    let (b, c): (int, int) = next_two(100i);
    print_number(b, c); 
    structs();
    order();
    matches();
    fors();
    strings();
    vectors();
	generics();
    println!( "North => {}", North as int );
}

fn print_number(x: int, y: int) {
    println!("print_number is: {}", x + y);
}

fn add_one(x: int) -> int {
	x + 1
}

fn tuples() {
	let x = (1i, "hello");
	let y: (int, &str) = (1, "hello");

	let a = (1i, 2i, 3i);
	let (b, c, d) = a;
	println!("a is {}", a);
	println!("b + c + d is {}", b + c + d);
	println!("x is {}", x);
	println!("y is {}", y);
}

fn next_two(x: int) -> (int, int) { (x + 1i, x + 2i) }

struct Point{
	x: int,
	y: int,
}

fn structs() {
	let mut point = Point { x: 0i, y: 0i };
	point.x = 5;
	println!("The point is at ({}, {})", point.x, point.y)
}

#[deriving(PartialEq)] enum Ordering {
	Less,
	Equal,
	Greater,
}

fn cmp(a: int, b: int) -> Ordering {
	if a < b { Less }
	else if a > b { Greater }
	else { Equal }
}

fn order() {
	let x = 5i;
	let y = 10i;

	let ordering = cmp(x, y);

	if ordering == Less {
        println!("less");
    } else if ordering == Greater {
        println!("greater");
    } else if ordering == Equal {
        println!("equal");
    }
}

fn matches() {
	let x = 5i;

	match x {
	    1 => println!("one"),
	    2 => println!("two"),
	    3 => println!("three"),
	    4 => println!("four"),
	    5 => println!("five"),
	    _ => println!("something else"),
	}
}

fn fors() {
	for x in range(0i, 10i) {
	    println!("{:d}", x);
	}
}

fn takes_slice(slice: &str) {
    println!("Got: {}", slice);
}

fn strings() {
	let mut s = "Hello".to_string();
	println!("{}", s);

	s.push_str(", world.");
	println!("{}", s);

	let x = 50005i;
	println!("{}", x.to_string().as_slice())

    let f = "Hello".to_string();
    takes_slice(f.as_slice());
}

fn vectors() {
	//Vector
	let mut nums_vec = vec![1i, 2i, 3i];
	nums_vec.push(4i);

	//Array
	let nums_arr = [1i, 2i, 3i];

	let slice = nums_vec.as_slice();
	println!("nums_vec.as_slice():{}", slice)

	for i in nums_vec.iter() {
		println!("{}", i)
	}

	let names = ["Graydon", "Brian", "Niko"];
	println!("The second name is: {}", names[1]);
}

fn generics() {
	let c = Circle {
        x: 0.0f64,
        y: 0.0f64,
        radius: 1.0f64,
    };

    let s = Square {
        x: 0.0f64,
        y: 0.0f64,
        side: 1.0f64,
    };

    print_area(c);
    print_area(s);
}

/*
fn inverse<T>(x: T) -> Result<T, String> {
	if x == 0.0 {return Err("x cannot be zero!".to_string()); }

	Ok(1.0 / x)
}
*/

trait HasArea {
    fn area(&self) -> f64;
}

struct Circle {
    x: f64,
    y: f64,
    radius: f64,
}

impl HasArea for Circle {
    fn area(&self) -> f64 {
        std::f64::consts::PI * (self.radius * self.radius)
    }
}

struct Square {
    x: f64,
    y: f64,
    side: f64,
}

impl HasArea for Square {
    fn area(&self) -> f64 {
        self.side * self.side
    }
}

fn print_area<T: HasArea>(shape: T) {
    println!("This shape has an area of {}", shape.area());
}