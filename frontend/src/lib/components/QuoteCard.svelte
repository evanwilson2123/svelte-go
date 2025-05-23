<script lang="ts">
    import type { DailyQuote } from "$lib/types/types";
    export let data: DailyQuote;
</script>

<div class="max-w-md mx-auto mt-8 p-6 bg-white dark:bg-gray-800 rounded-2xl shadow-lg">
    <h1 class="text-2xl font-bold text-gray-900 dark:text-white mb-4 text-center">
      Daily Stoic Quote
    </h1>
  
    {#if data}
      <p class="text-sm uppercase text-gray-500 dark:text-gray-400 mb-2 text-center">
        — {data.author}
      </p>
      <blockquote class="text-lg italic text-gray-800 dark:text-gray-200 mb-4 px-4 border-l-4 border-white">
        “{data.quote}”
      </blockquote>
      <p class="text-xs text-gray-400 dark:text-gray-600 text-right">
        {data.dateWritten}
      </p>
    {:else}
      <p class="text-center text-gray-500 dark:text-gray-400">No Data Available</p>
    {/if}
  </div>