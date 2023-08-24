<script lang="ts">
	import { getAPIEndpoint } from '$lib';
	import type { EventHandler } from 'svelte/elements';
	import { createMutation } from '@tanstack/svelte-query';

	function getFormItem<T extends Element>(form: HTMLFormElement, name: string) {
		return form.elements.namedItem(name) as T | null;
	}

	const mutation = createMutation({
		mutationKey: ['auth:login'],
		mutationFn: async (payload: { email?: string; password?: string }) => {
			const resp = await fetch(`${getAPIEndpoint()}/v1/auth/login`, {
				method: 'POST',
				body: JSON.stringify(payload),
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json'
				}
			});

			const data = await resp.json();
			return data;
		}
	});

	const login: EventHandler<SubmitEvent, HTMLFormElement> = (e) => {
		const form = e.target as HTMLFormElement;
		const email = getFormItem<HTMLInputElement>(form, 'email')?.value;
		const password = getFormItem<HTMLInputElement>(form, 'password')?.value;
		$mutation.mutate({ email, password });
	};
</script>

<svelte:head>
	<title>sora.mon0 | Authenticate</title>
	<meta name="description" content="sora.mon0's full stack portfolio website" />
</svelte:head>

<main class="h-screen p-4 sm:p-6 flex items-center justify-center">
	<section class="space-y-4 sm:space-y-8">
		<h1 class="text-4xl">Admin sign in</h1>
		<form class="block" on:submit|preventDefault={login}>
			<fieldset class="space-y-4" disabled={$mutation.isLoading}>
				<div class="flex items-left gap-2 flex-col">
					<label for="email">Email</label>
					<input type="email" id="email" required />
				</div>
				<div class="flex items-left gap-2 flex-col">
					<label for="password">Password</label>
					<input type="password" id="password" required />
				</div>
				<button type="submit">
					{#if $mutation.isLoading}
						<span>Loading...</span>
					{:else}
						<span>Login</span>
					{/if}
				</button>
			</fieldset>
		</form>
	</section>
</main>
