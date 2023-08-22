import { PUBLIC_API_ENDPOINT } from '$env/static/public';

function checkEnv(env: string | undefined, name: string) {
	if (!env) {
		throw new Error(`Please define the ${name} environment variable inside .env`);
	}

	return env;
}

const isDev = process.env.NODE_ENV === 'development';

export function getAPIEndpoint() {
	if (isDev) {
		return checkEnv(PUBLIC_API_ENDPOINT, 'PUBLIC_API_ENDPOINT');
	}

	return '/api';
}
