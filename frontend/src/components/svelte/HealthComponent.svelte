<script>
  // Ping the server to check if it's up and running, localhost:3001/api/health
  const healthCheck = async () => {
    console.log("Checking health...");
    try {
      const response = await fetch("http://localhost:3001/api/health");
      const text = await response.text();
      const data = JSON.parse(text);
      console.log(data);
      return typeof data.status === "object"
        ? JSON.stringify(data.status)
        : data.status;
    } catch (error) {
      console.error("Error fetching health check:", error);
      return "Failed to fetch health check";
    }
  };
</script>

<main>
  <h1>Health Check</h1>
  {#await healthCheck()}
    <p class="font-bold">Checking...</p>
  {:then status}
    <p><span class="font-bold">Status:</span> {status}</p>
  {/await}
</main>
