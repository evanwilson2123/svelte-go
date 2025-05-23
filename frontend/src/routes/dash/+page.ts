import { fetcher } from '$lib/utils/fetcher';

export async function load() {
	try {
		const user = await fetcher<any>('/users');
		console.log(user);
		console.log('name 2', user.name);
		return { name: user.name };
	} catch (error: any) {
		console.error(error);
		return { user: null };
	}
}
