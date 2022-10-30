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

      <main className="h-screen">
        <HeaderBanner />

        <section className="h-5/6 mx-auto max-w-8xl flex items-center justify-center">
          <h1 className="text-3xl leading-snug font-display font-bold text-slate-800 sm:text-4xl md:text-5xl lg:text-6xl">
            Fullstack <span className="text-indigo-600">Developer</span>
          </h1>
        </section>
      </main>
    </>
  );
}

export default HomePage;
