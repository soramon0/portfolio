import { getAPIEndpoint } from '$lib';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const resp = await fetch(`${getAPIEndpoint()}/healthz`);
	const data = (await resp.json()) as { ok: boolean };

	return {
		data
	};
};
