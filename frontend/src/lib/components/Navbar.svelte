<script lang="ts">
    let isOpen = false;
  </script>
  
  <nav class="bg-white dark:bg-gray-900 shadow outline-4 outline-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16 items-center">
        <!-- Logo / Site Title -->
        <div class="text-xl font-semibold text-gray-800 dark:text-white">
          <a href="/">Stoic Journal</a>
        </div>
  
        <!-- Mobile Menu Button -->
        <div class="md:hidden">
          <button
            class="text-gray-700 dark:text-gray-300 focus:outline-none"
            on:click={() => (isOpen = !isOpen)}
          >
            {#if isOpen}
              ✖
            {:else}
              ☰
            {/if}
          </button>
        </div>
  
        <!-- Links (Desktop) -->
        <div class="hidden md:flex space-x-6">
          <a href="/" class="text-gray-700 dark:text-gray-300 hover:underline">Home</a>
          <a href="/history" class="text-gray-700 dark:text-gray-300 hover:underline">History</a>
          <a href="/about" class="text-gray-700 dark:text-gray-300 hover:underline">About</a>
        </div>
      </div>
    </div>
  
    <!-- Links (Mobile Dropdown) -->
    {#if isOpen}
      <div class="md:hidden px-4 pb-4 space-y-2">
        <a href="/" class="block text-gray-700 dark:text-gray-300">Home</a>
        <a href="/history" class="block text-gray-700 dark:text-gray-300">History</a>
        <a href="/about" class="block text-gray-700 dark:text-gray-300">About</a>
      </div>
    {/if}
  </nav>
  