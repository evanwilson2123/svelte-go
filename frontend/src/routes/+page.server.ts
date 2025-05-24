import type { DailyQuote } from '$lib/types/types';
import { fetcher } from '$lib/utils/fetcher';

export async function load() {
	const quote = await fetcher<DailyQuote>('/');
	return { quote };
}
