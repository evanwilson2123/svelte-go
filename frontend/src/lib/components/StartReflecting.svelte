<script lang="ts">
    import { goto } from '$app/navigation';
    
    let name = '';
    let email = '';
    let password = '';
    let error = '';

    async function handleSubmit() {
        const res = await fetch(`${import.meta.env.VITE_BASE_API}/start-reflecting`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, password })
        });

        if (res.ok) {
            goto('/dash');
        } else {
            const result = await res.json();
            error = result.error || 'Signup failed';
        }
    }
</script>


<div class="max-w-md mx-auto mt-12 p-6 bg-gray-800 text-white rounded-2xl shadow-lg">
    <h1 class="text-2xl font-bold mb-4 text-center">Create an Account</h1>
  
    {#if error}
      <p class="text-red-400 mb-2">{error}</p>
    {/if}
  
    <form on:submit|preventDefault={handleSubmit} class="space-y-4">
      <input class="w-full p-2 rounded bg-gray-900" bind:value={name} placeholder="Name" required />
      <input class="w-full p-2 rounded bg-gray-900" type="email" bind:value={email} placeholder="Email" required />
      <input class="w-full p-2 rounded bg-gray-900" type="password" bind:value={password} placeholder="Password" required />
  
      <button class="w-full py-2 bg-white hover:bg-gray-300 rounded text-gray-800 font-semibold hover:cursor-pointer">
        Sign Up
      </button>
    </form>
  </div>
  