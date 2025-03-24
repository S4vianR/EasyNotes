<script>
  // Declare variable for notes
  let notes = [];

  // Function to fetch notes from the backend
  const notesCallback = () => {
    // Call to localhost:3001/api/notes
    fetch("http://localhost:3001/api/notes")
      .then((response) => response.json())
      .then((data) => {
        // Set the notes variable to the data from the backend
        notes = data;
      });
  };

  $: notesCallback();
</script>

<section class="px-2 py-4">
  <h1 class="text-2xl font-bold text-center">Notes</h1>
  <ul class="list-none grid grid-cols-4 gap-4">
    <!-- Loop through the notes and display them -->
    {#each notes as note}
      <li class="p-1 min-w-fit w-full border-2 rounded-lg">
        <p class="text-lg font-bold">{note.title}</p>
        <p class="text-sm text-gray-400">{note.content}</p>
      </li>
    {/each}
  </ul>
</section>
