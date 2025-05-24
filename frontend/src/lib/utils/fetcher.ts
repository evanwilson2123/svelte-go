export async function fetcher<T>(url?: string): Promise<T | null> {
	try {
		const res = await fetch(`${import.meta.env.VITE_BASE_API}${url}`);
		return await res.json();
	} catch (error: any) {
		console.error(error);
		return null;
	}
}
