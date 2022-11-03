import { MegaphoneIcon } from '@heroicons/react/24/outline';

import IconGithub from './IconGithub';

function HeaderBanner() {
  return (
    <div className="bg-indigo-600">
      <div className="mx-auto max-w-8xl py-3 px-3 sm:px-6 lg:px-8">
        <div className="flex flex-wrap items-center justify-between">
          <div className="flex w-0 flex-1 items-center">
            <span className="flex rounded-lg bg-indigo-800 p-2">
              <MegaphoneIcon
                className="h-6 w-6 text-white"
                aria-hidden="true"
              />
            </span>

            <p className="ml-3 truncate font-medium text-white">
              <span className="md:hidden">launching new portfolio!</span>
              <span className="hidden md:inline">
                Big news! sora.mon0 portfolio is about to be launched
              </span>
            </p>
          </div>
          <a href="https://github.com/soramon0" target="_blank">
            <IconGithub className="w-6 h-6 text-white" />
            <span className="sr-only">View Github profile</span>
          </a>
        </div>
      </div>
    </div>
  );
}

export default HeaderBanner;
