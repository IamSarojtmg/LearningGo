// server.js (Node.js with Express)
const express = require('express');
const app = express();

// Example time series data
const timeSeriesData = {
    labels: ['2024-11-01', '2024-11-02', '2024-11-03', '2024-11-04', '2024-11-05'],
    datasets: [{
        label: 'Data Points',
        data: [5, 10, 15, 20, 25],
        borderColor: 'rgba(75, 192, 192, 1)',
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        fill: true,
        tension: 0.4
    }]
};

// Serve static files like HTML, CSS, and JS
app.use(express.static('public'));

// Endpoint to fetch time series data
app.get('/data', (req, res) => {
    res.json(timeSeriesData); // Return time series data as JSON
});

// Start the server
app.listen(3000, () => {
    console.log('Server is running on http://localhost:3000');
});
