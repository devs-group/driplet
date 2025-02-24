import * as React from "react";

const Footer = React.forwardRef<HTMLDivElement, React.HTMLAttributes<HTMLDivElement>>(
  ({ className, ...props }, ref) => {
    return (
      <footer className="flex gap-6 flex-wrap items-center justify-center pb-6" ref={ref} {...props}>
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
    );
  }
);

Footer.displayName = "Footer";

export default Footer;
