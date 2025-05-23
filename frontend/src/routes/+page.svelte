<script lang="ts">
    import Hero from '$lib/components/Hero.svelte';
    import QuoteCard from '$lib/components/QuoteCard.svelte';
    import type { DailyQuote } from '$lib/types/types';
  
    export let data: { quote: DailyQuote };
  </script>
  
  <!-- Page container -->
  <div class="min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-white">
    <!-- Main content wrapper -->
    <main class="max-w-4xl mx-auto px-4 py-12 space-y-10">
      
      <!-- Daily Quote section -->
      <QuoteCard data={data.quote} />
        <!-- Hero section -->
      <Hero />
  
      
    </main>
  </div>
  