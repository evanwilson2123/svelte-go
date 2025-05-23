<script lang="ts">
    export let name = 'User';
  </script>
  
  <div class="max-w-3xl mx-auto mt-12 p-6 bg-gray-900 text-white rounded-2xl shadow-lg">
    <h1 class="text-3xl font-bold mb-6 text-center">Welcome back, {name} 👋</h1>
  
    <div class="grid gap-4 sm:grid-cols-2">
      <div class="p-4 bg-gray-800 rounded-xl shadow">
        <h2 class="text-lg font-semibold mb-2">Today's Quote</h2>
        <p class="text-sm text-gray-300">"You have power over your mind — not outside events. Realize this, and you will find strength."</p>
      </div>
  
      <div class="p-4 bg-gray-800 rounded-xl shadow">
        <h2 class="text-lg font-semibold mb-2">Reflections</h2>
        <p class="text-sm text-gray-300">You haven't submitted a reflection today yet.</p>
      </div>
  
      <div class="p-4 bg-gray-800 rounded-xl shadow sm:col-span-2">
        <h2 class="text-lg font-semibold mb-2">Progress</h2>
        <p class="text-sm text-gray-300">You've reflected 4 days in a row — keep going!</p>
      </div>
    </div>
  </div>
  