import enUS from './en-us';
import zhCN from './zh-cn';
import { createI18n } from 'vue-i18n';

const messages = {
  en_US: {
    ...enUS,
  },
  zh_CN: {
    ...zhCN,
  },
};

const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  silentTranslationWarn: true,
  locale: 'zh_CN',
  fallbackLocale: 'zh_CN',
  messages,
});

export default i18n;
