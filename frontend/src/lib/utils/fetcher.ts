import type { DailyQuote } from '$lib/types/types';

export async function fetcher(url?: string): Promise<DailyQuote | null> {
	try {
		const res = await fetch(`${import.meta.env.VITE_BASE_API}${url}`);
		return await res.json();
	} catch (error: any) {
		console.error(error);
		return null;
	}
}
