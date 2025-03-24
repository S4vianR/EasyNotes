<script>
  // Ping the server to check if it's up and running, localhost:3001/api/health
  const healthCheck = async () => {
    console.log("Checking health...");
    try {
      const response = await fetch("http://localhost:3001/api/health");
      const text = await response.text();
      const data = JSON.parse(text);
      console.log(data);
      return typeof data.status === 'object' ? JSON.stringify(data.status) : data.status;
    } catch (error) {
      console.error("Error fetching health check:", error);
      return "Failed to fetch health check";
    }
  };
</script>

<main>
  <h1>Health Check</h1>
  <p>Checking the health of the server...</p>
  {#await healthCheck()}
    <p>Checking...</p>
  {:then status}
    <p>Status: {status}</p>
  {:catch error}
    <p>Error: {error}</p>
  {/await}
</main>