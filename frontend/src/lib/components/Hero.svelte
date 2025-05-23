<script lang="ts">
    export let title: string = "Welcome to Stoic Journal";
    export let subtitle: string = "Daily reflections inspired by timeless philosophy.";
  </script>
  
  <div class="max-w-md mx-auto mt-12 p-6 bg-white dark:bg-gray-800 rounded-2xl shadow-lg text-center">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-4">
      {title}
    </h1>
  
    <p class="text-base text-gray-600 dark:text-gray-300 mb-4">
      {subtitle}
    </p>
  
    <a
      href="/start-reflecting"
      class="inline-block mt-4 px-4 py-2 text-sm font-semibold text-gray-800 bg-white hover:bg-gray-400 rounded-full shadow-md transition-colors"
    >
      Start Reflecting
    </a>
  </div>
  