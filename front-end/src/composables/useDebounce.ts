export function useDebounce() {
  let timer: NodeJS.Timeout;

  const debounce = <T extends (...args: any[]) => void>(
    fn: T,
    delay: number
  ) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      fn();
    }, delay);
  };

  return {
    debounce,
  };
}
