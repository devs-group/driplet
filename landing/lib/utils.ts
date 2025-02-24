import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

/**
 * A utility function for conditionally joining class names together.
 * Uses clsx for conditional class merging and twMerge to handle Tailwind CSS conflicts.
 *
 * @param inputs - Class names or conditional class objects to merge
 * @returns A string of merged class names
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
