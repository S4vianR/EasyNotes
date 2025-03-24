<script>
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
        alert("Note created successfully!");
        document.getElementById("title").value = "";
        document.getElementById("note").value = "";
      } else {
        alert("Failed to create note");
      }
    } catch (error) {
      console.error("Error creating note:", error);
      alert("Failed to create note");
    } finally {
      // Call the notesCallback function to update the notes
      notesCallback();
    }
  };
</script>

<section class="w-full">
  <form
    class="p-1 flex flex-col items-center justify-center gap-2"
    onsubmit="{createNote}"
  >
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
</section>
