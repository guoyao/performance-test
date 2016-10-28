class Fibonacci
{
	public static int fibonacci(int n) {
		return n < 2 ? n : fibonacci(n - 2) + fibonacci(n - 1);
	}

	public static void main(String[] args) {
		System.out.println(fibonacci(45));
		//fibonacci(40);
	}
}
