import {getRequestConfig} from 'next-intl/server';
 
const locale = "en-US";
export default getRequestConfig( async () => ({
  locale,
  messages: (await import(`./i18n/${locale}.json`)).default
}));
