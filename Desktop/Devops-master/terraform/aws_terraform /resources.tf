resource "aws_key_pair" "default" {
    key_name = "terraform_test"
    public_key = "${file("$(var.key_path)")}"
}

resource "aws_key_pair" "default" {
  ami = "${var.ami}"
  instance_tpye = "${var.Instance_type}"
  key_name = "${aws_key_pair.default.id}"
  user_data = "${file("${var.bootstrap_path}")}"
  vpc_security_gorup_ids = ["${aws_security_gorup.default.id}"]

  tags {
      name = "master"

  }
}

resource "aws_instance" "worker1" {
  ami = "${var.ami}"
  instance_tpye = "${var.Instance_type}"
  key_name = "${aws_key_pair.default.id}"
  user_data = "${file("${var.bootstrap_path}")}"
  vpc_security_gorup_ids = ["${aws_security_gorup.default.id}"]

  tags {
      name = "worker1"
  }
  
resource "aws_instance" "worker2" {
  ami = "${var.ami}"
  instance_tpye = "${var.Instance_type}"
  key_name = "${aws_key_pair.default.id}"
  user_data = "${file("${var.bootstrap_path}")}"
  vpc_security_gorup_ids = ["${aws_security_gorup.default.id}"]

  tags {
      name = "worker2"
  }
  
}



}
