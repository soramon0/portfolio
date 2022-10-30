import Head from 'next/head';

import { HeaderBanner } from '@/components';

function HomePage() {
  return (
    <>
      <Head>
        <title>Home</title>
        <meta
          name="description"
          content="sora.mon0's full stack portfolio website"
        />
      </Head>

      <HeaderBanner />
      <main className="h-screen mx-auto max-w-8xl flex items-center justify-center">
        <h1 className="text-4xl leading-snug font-display font-bold text-slate-800 md:text-5xl lg:text-6xl">
          Fullstack <span className="text-indigo-600">Developer!</span>
        </h1>
      </main>
    </>
  );
}

export default HomePage;
