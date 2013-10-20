def fibonacci(n)
  return n < 2 ? n : fibonacci(n - 2) + fibonacci(n - 1)
end

fibonacci(40);
# puts fibonacci(40)
