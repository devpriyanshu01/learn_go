// server.js

const express = require('express');
const app = express();
const PORT = 3000;

// Middleware to parse JSON bodies
app.use(express.json());

// Sample GET route
app.get('/', (req, res) => {
  res.send('Hello, world!');
});

// GET route with a parameter
app.get('/hello/:name', (req, res) => {
  const { name } = req.params;
  res.send(`Hello, ${name}!`);
});

// POST route to echo back data
app.post('/echo', (req, res) => {
  const data = req.body;
  res.json({ received: data });
});

// POST route to simulate user creation
app.post('/users', (req, res) => {
  const { username, email } = req.body;
  // Here you'd normally save the user to a database
  res.status(201).json({ message: 'User created', user: { username, email } });
});

// Start server
app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});
