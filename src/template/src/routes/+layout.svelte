<script>
	import '../app.css';
	import { browser } from '$app/environment';
	import { afterUpdate, beforeUpdate } from 'svelte';
	// @ts-ignore
	import { useAckeeSapper } from 'svelte-ackee';
	import { PUBLIC_ACKE_SERVER, PUBLIC_ACKE_DOMAIN_ID } from '$env/static/public';

	if (browser) {
		useAckeeSapper(
			beforeUpdate,
			afterUpdate,
			{
				server: PUBLIC_ACKE_SERVER,
				domainId: PUBLIC_ACKE_DOMAIN_ID
			},
			{
				ignoreLocalhost: false
			}
		);
	}
</script>

<slot />
