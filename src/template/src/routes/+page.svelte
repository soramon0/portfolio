<script lang="ts">
	import type { PageData } from './$types';
		
	 export let data: PageData;
	 console.log('home page data:', data)
</script>

<head>
	<title>sora.mon0 | Home</title>
	<meta name="description" content="sora.mon0's full stack portfolio website" />
</head>

<main class="h-screen">
	<div class="bg-indigo-600">
		<div class="mx-auto max-w-8xl py-3 px-3 sm:px-6 lg:px-8">
			<div class="flex flex-wrap items-center justify-between">
				<div class="flex w-0 flex-1 items-center">
					<span class="flex rounded-lg bg-indigo-800 p-2">
						<svg
							class="h-6 w-6 text-white"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M10.34 15.84c-.688-.06-1.386-.09-2.09-.09H7.5a4.5 4.5 0 110-9h.75c.704 0 1.402-.03 2.09-.09m0 9.18c.253.962.584 1.892.985 2.783.247.55.06 1.21-.463 1.511l-.657.38c-.551.318-1.26.117-1.527-.461a20.845 20.845 0 01-1.44-4.282m3.102.069a18.03 18.03 0 01-.59-4.59c0-1.586.205-3.124.59-4.59m0 9.18a23.848 23.848 0 018.835 2.535M10.34 6.66a23.847 23.847 0 008.835-2.535m0 0A23.74 23.74 0 0018.795 3m.38 1.125a23.91 23.91 0 011.014 5.395m-1.014 8.855c-.118.38-.245.754-.38 1.125m.38-1.125a23.91 23.91 0 001.014-5.395m0-3.46c.495.413.811 1.035.811 1.73 0 .695-.316 1.317-.811 1.73m0-3.46a24.347 24.347 0 010 3.46"
							/>
						</svg>
					</span>

					<p class="ml-3 truncate font-medium text-white">
						<span class="md:hidden">launching new portfolio!</span>
						<span class="hidden md:inline">
							Big news! sora.mon0 portfolio is about to be launched
						</span>
					</p>
				</div>
				<a href="https://github.com/soramon0" target="_blank" rel="noreferrer">
					<svg
						fill="none"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						aria-hidden="true"
						class="w-6 h-6 text-white"
					>
						<path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M12 0C5.37 0 0 5.50583 0 12.3035C0 17.7478 3.435 22.3463 8.205 23.9765C8.805 24.0842 9.03 23.715 9.03 23.3921C9.03 23.0999 9.015 22.131 9.015 21.1005C6 21.6696 5.22 20.347 4.98 19.6549C4.845 19.3012 4.26 18.2092 3.75 17.917C3.33 17.6863 2.73 17.1173 3.735 17.1019C4.68 17.0865 5.355 17.9939 5.58 18.363C6.66 20.2239 8.385 19.701 9.075 19.3781C9.18 18.5783 9.495 18.04 9.84 17.7325C7.17 17.4249 4.38 16.3637 4.38 11.6576C4.38 10.3196 4.845 9.21227 5.61 8.35102C5.49 8.04343 5.07 6.78232 5.73 5.09058C5.73 5.09058 6.735 4.76762 9.03 6.3517C9.99 6.07487 11.01 5.93645 12.03 5.93645C13.05 5.93645 14.07 6.07487 15.03 6.3517C17.325 4.75224 18.33 5.09058 18.33 5.09058C18.99 6.78232 18.57 8.04343 18.45 8.35102C19.215 9.21227 19.68 10.3042 19.68 11.6576C19.68 16.3791 16.875 17.4249 14.205 17.7325C14.64 18.1169 15.015 18.8552 15.015 20.0086C15.015 21.6542 15 22.9768 15 23.3921C15 23.715 15.225 24.0995 15.825 23.9765C18.2072 23.1519 20.2773 21.5822 21.7438 19.4882C23.2103 17.3942 23.9994 14.8814 24 12.3035C24 5.50583 18.63 0 12 0Z"
							fill="currentColor"
						/>
					</svg>
					<span class="sr-only">View Github profile</span>
				</a>
			</div>
		</div>
	</div>

	<section class="relative px-4 h-5/6 mx-auto max-w-8xl flex items-center justify-center">
		<h1
			class="text-3xl leading-snug font-display font-bold text-slate-800 sm:text-4xl md:text-5xl lg:text-6xl"
		>
			Fullstack <span class="text-indigo-600">Developer</span>
		</h1>
	</section>
</main>
