import Link from "next/link";
import Image from "next/image";

export default function Home() {
  return (
    <div className="grid grid-rows-[1fr_auto] min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-16 items-center justify-center max-w-4xl mx-auto text-center pt-24">
        {/* Hero Section */}
        <div className="space-y-6">
          <div className="inline-block mb-4 px-4 py-2 bg-blue-50 rounded-full dark:bg-blue-900/30">
            <span className="text-sm font-medium text-white font-[family-name:var(--font-geist-mono)]">
              Coming Soon — Join Waitlist
            </span>
          </div>

          <h1 className="text-4xl sm:text-6xl font-bold tracking-tight">
            Get paid to browse
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-blue-500 to-indigo-600 ml-2">
              with Driplet
            </span>
          </h1>

          <p className="mt-6 text-lg sm:text-xl text-gray-600 dark:text-gray-300 max-w-2xl mx-auto">
            Browse the web, share anonymous data, and earn $DRIPL tokens
            directly to your wallet.
          </p>

          <div className="flex gap-4 items-center justify-center mt-10 flex-col sm:flex-row">
            <Link
              href="/waitlist"
              className="rounded-full border border-solid border-transparent transition-colors flex items-center justify-center bg-white text-black gap-2 hover:bg-gray-200 text-sm sm:text-base h-12 sm:h-14 px-8 sm:px-10 font-medium"
            >
              Join Waitlist
            </Link>
            <Link
              href="/learn-more"
              className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-12 sm:h-14 px-8 sm:px-10 font-medium"
            >
              Learn More
            </Link>
          </div>
        </div>

        {/* Browser Extension Mockup */}
        <div className="relative w-full max-w-3xl mx-auto mt-10 mb-10">
          <div className="aspect-[16/9] relative rounded-lg overflow-hidden border border-gray-200 dark:border-gray-800 shadow-2xl bg-white dark:bg-gray-900">
            <div className="absolute inset-0 bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-blue-950/30 dark:to-indigo-950/30 flex items-center justify-center">
              <div className="text-center p-8 max-w-md">
                <Image
                  src="/extension.png"
                  alt="Browser Extension Mockup"
                  width={122}
                  height={120}
                  className="w-full h-full object-contain"
                />
              </div>
            </div>
          </div>

          {/* Browser Chrome */}
          <div className="absolute -top-2 inset-x-0 h-10 bg-gray-100 dark:bg-gray-800 rounded-t-lg border-b border-gray-200 dark:border-gray-700 flex items-center px-4">
            <div className="flex space-x-2">
              <div className="w-3 h-3 rounded-full bg-gray-300 dark:bg-gray-600"></div>
              <div className="w-3 h-3 rounded-full bg-gray-300 dark:bg-gray-600"></div>
              <div className="w-3 h-3 rounded-full bg-gray-300 dark:bg-gray-600"></div>
            </div>
            <div className="mx-auto bg-white dark:bg-gray-700 rounded-full h-6 w-1/2 flex items-center justify-center">
              <span className="text-xs text-gray-500 dark:text-gray-400 font-[family-name:var(--font-geist-mono)]">
                driplet.io
              </span>
            </div>
          </div>
        </div>

        {/* Key Features */}
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-8 w-full max-w-3xl">
          <div className="flex flex-col items-center text-center p-4">
            <div className="w-12 h-12 rounded-full bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center mb-4">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
                className="w-6 h-6 text-blue-600 dark:text-blue-400"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M9.813 15.904L9 18.75l-.813-2.846a4.5 4.5 0 0 0-3.09-3.09L2.25 12l2.846-.813a4.5 4.5 0 0 0 3.09-3.09L9 5.25l.813 2.846a4.5 4.5 0 0 0 3.09 3.09L15.75 12l-2.846.813a4.5 4.5 0 0 0-3.09 3.09ZM18.259 8.715 18 9.75l-.259-1.035a3.375 3.375 0 0 0-2.455-2.456L14.25 6l1.036-.259a3.375 3.375 0 0 0 2.455-2.456L18 2.25l.259 1.035a3.375 3.375 0 0 0 2.456 2.456L21.75 6l-1.035.259a3.375 3.375 0 0 0-2.456 2.456ZM16.894 20.567 16.5 21.75l-.394-1.183a2.25 2.25 0 0 0-1.423-1.423L13.5 18.75l1.183-.394a2.25 2.25 0 0 0 1.423-1.423l.394-1.183.394 1.183a2.25 2.25 0 0 0 1.423 1.423l1.183.394-1.183.394a2.25 2.25 0 0 0-1.423 1.423Z"
                />
              </svg>
            </div>
            <h3 className="text-lg font-semibold mb-2 font-[family-name:var(--font-geist-sans)]">
              Private & Secure
            </h3>
            <p className="text-sm text-gray-600 dark:text-gray-400 font-[family-name:var(--font-geist-mono)]">
              Your personal data stays private. Share only anonymized browsing
              patterns.
            </p>
          </div>

          <div className="flex flex-col items-center text-center p-4">
            <div className="w-12 h-12 rounded-full bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center mb-4">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
                className="w-6 h-6 text-blue-600 dark:text-blue-400"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M12 6v12m-3-2.818.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                />
              </svg>
            </div>
            <h3 className="text-lg font-semibold mb-2 font-[family-name:var(--font-geist-sans)]">
              Earn $DRIPL
            </h3>
            <p className="text-sm text-gray-600 dark:text-gray-400 font-[family-name:var(--font-geist-mono)]">
              Get rewarded with $DRIPL tokens directly to your connected wallet.
            </p>
          </div>

          <div className="flex flex-col items-center text-center p-4">
            <div className="w-12 h-12 rounded-full bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center mb-4">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
                className="w-6 h-6 text-blue-600 dark:text-blue-400"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M3.75 13.5l10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75z"
                />
              </svg>
            </div>
            <h3 className="text-lg font-semibold mb-2 font-[family-name:var(--font-geist-sans)]">
              Seamless
            </h3>
            <p className="text-sm text-gray-600 dark:text-gray-400 font-[family-name:var(--font-geist-mono)]">
              Install once and browse normally. Earnings accumulate
              automatically.
            </p>
          </div>
        </div>
      </main>

      <footer className="flex gap-6 flex-wrap items-center justify-center pb-6">
        <a
          className="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 transition-colors font-[family-name:var(--font-geist-mono)]"
          href="https://twitter.com/driplet"
          target="_blank"
          rel="noopener noreferrer"
        >
          Twitter
        </a>
        <a
          className="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 transition-colors font-[family-name:var(--font-geist-mono)]"
          href="https://github.com/devs-group/driplet"
          target="_blank"
          rel="noopener noreferrer"
        >
          GitHub
        </a>
        <a
          className="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 transition-colors font-[family-name:var(--font-geist-mono)]"
          href="/privacy"
        >
          Privacy
        </a>
        <a
          className="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 transition-colors font-[family-name:var(--font-geist-mono)]"
          href="/terms"
        >
          Terms
        </a>
      </footer>
    </div>
  );
}
