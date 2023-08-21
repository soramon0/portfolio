import { FC, PropsWithChildren, ReactNode } from 'react';
import { AppProps } from 'next/app';
import { Analytics } from '@vercel/analytics/react';

import '@/styles/global.css';

interface AppWithLayoutProps extends AppProps {
  Component: AppProps['Component'] & {
    Layout: FC<PropsWithChildren>;
  };
}

const Noop = ({ children }: { children: ReactNode }) => children;

export default function MyApp({ Component, pageProps }: AppWithLayoutProps) {
  const Layout = Component.Layout || Noop;

  return (
    <Layout>
      <Component {...pageProps} />
      <Analytics />
    </Layout>
  );
}
