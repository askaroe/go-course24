<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>

<body>

  <h1>Chunin Exam API</h1>


  <h2>Description</h2>

  <p>The Chunin Exam API is a simple API for managing data related to Chunin exams in the world of Naruto. It provides endpoints to retrieve information about genins participating in the exams.</p>

  <h2>Table of Contents</h2>

  <ul>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#endpoints">Endpoints</a></li>
  </ul>

  <h2>Installation</h2>

  <p>To use the Chunin Exam API, follow these installation steps:</p>

  <pre><code># Clone the repository
git clone https://github.com/askaroe/go-course24/tree/main/tsis1-chunin-exam-api

# Change to the project directory
cd tsis1-chunin-exam-api

# Build the project
go build

# Run the API server
go run main.go
</code></pre>

  <h2>Usage</h2>

  <p>Once the API is running, you can interact with it using HTTP requests. See the <a href="#endpoints">Endpoints</a> section for available routes and examples.</p>

  <h2>Endpoints</h2>

  <ul>
    <li><strong>Health Check:</strong>
      <ul>
        <li><code>GET /health-check</code>: Check the health of the API.</li>
      </ul>
    </li>
    <li><strong>List Genins:</strong>
      <ul>
        <li><code>GET /genins</code>: Retrieve a list of genins participating in the Chunin exams.</li>
      </ul>
    </li>
    <li><strong>Get Genin by Name:</strong>
      <ul>
        <li><code>GET /genins/{name}</code>: Retrieve information about a specific genin by name.</li>
      </ul>
    </li>
  </ul>

</body>

</html>
