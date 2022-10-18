Bun.serve({
  port: 8080,
  fetch(request, response) {
    return new Response(JSON.stringify({ message: "Hello World" }), {
      status: 200,
      headers: { "Content-Type": "application/json" },
    });
  },
});
