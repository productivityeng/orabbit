import {getRequestConfig} from 'next-intl/server';
 
const locale = "pt-BR";
export default getRequestConfig( async () => ({
  messages: (await import(`./i18n/${locale}.json`)).default
}));
