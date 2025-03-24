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

<section class="w-full border-t">
  <h1>Notes</h1>
  <p>Here are some notes:</p>
  <ul class="list-none flex flex-row justify-start items-center gap-2">
    <!-- Loop through the notes and display them -->
    {#each notes as note}
      <li class="min-w-fit w-3/12 border">
        <p class="text-lg font-bold">{note.title}</p>
        <p class="text-sm text-gray-400">{note.content}</p>
      </li>
    {/each}
  </ul>
</section>
