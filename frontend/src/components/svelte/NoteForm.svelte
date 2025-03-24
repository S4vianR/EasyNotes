<script>
  let message = "";
  let noteCreatedSuccessfully = false;

  const createNote = async (event) => {
    event.preventDefault();
    const title = document.getElementById("title").value;
    const note = document.getElementById("note").value;
    const data = { title, content: note };
    try {
      const response = await fetch("http://localhost:3001/api/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (response.ok) {
        message = "Note created successfully!";
        noteCreatedSuccessfully = true;
        document.getElementById("title").value = "";
        document.getElementById("note").value = "";

        // Wait 300ms before resetting the message
        setTimeout(() => {
          // Bring the user back to the notes page
          window.location.href = "/";
        }, 500);
      } else {
        message = "Failed to create note";
        noteCreatedSuccessfully = false;
      }
    } catch (error) {
      console.error("Error creating note:", error);
      alert("Failed to create note");
    }
  };
</script>

<section class="p-5 flex flex-col items-center justify-center gap-2">
  <form
    class="p-1 w-full flex flex-col items-center justify-center gap-2"
    onsubmit={createNote}
  >
    <h1 class="text-2xl font-bold text-center">Create a Note!</h1>
    <label for="title" class="font-bold"> Title </label>
    <input
      id="title"
      name="title"
      type="text"
      placeholder="Enter the title..."
      class="p-1 w-4/12 rounded-md border-2 border-gray-500 focus:outline focus:ring focus:border-blue-300"
      required
    />
    <label for="content" class="font-bold"> Content </label>
    <textarea
      id="note"
      name="note"
      rows="4"
      placeholder="Write the note here..."
      class="p-1 w-4/12 rounded-md border-2 border-gray-500 focus:outline focus:ring focus:border-blue-300 resize-none"
    ></textarea>
    <button
      type="submit"
      class="w-24 h-10 bg-blue-500 text-white py-2 px-4 rounded cursor-pointer hover:bg-blue-600"
    >
      Save</button
    >
  </form>
  <p>
    <!-- Here where the message of success or failure goes -->
    {#if noteCreatedSuccessfully}
      <span class="text-green-500 font-bold">{message}</span>
    {:else}
      <span class="text-red-500 font-bold">{message}</span>
    {/if}
  </p>
</section>
