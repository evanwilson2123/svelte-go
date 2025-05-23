import { fetcher } from '$lib/utils/fetcher';

export async function load() {
	const quote = await fetcher('/');
	return { quote };
}
