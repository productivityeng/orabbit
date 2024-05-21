"use client";

import { useTranslations } from "next-intl";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";

export enum AVAILABLE_LANGUAGES {
    ENGLISH = 'en-US',
    PORTUGUESE = 'pt-BR'
}

export const LanguageSelector = () => {
    const t = useTranslations();
    const router = useRouter();

    const handleChange = (language: AVAILABLE_LANGUAGES) => {
        localStorage.setItem('language', language);
        router.refresh();
        toast.success(t('Common.LanguageChanged'));
    }

    return <Select onValueChange={handleChange} value={localStorage.getItem('language')?? AVAILABLE_LANGUAGES.ENGLISH}>
    <SelectTrigger >
      <SelectValue placeholder={t('Common.Language')}  />
    </SelectTrigger>
    <SelectContent >
      <SelectItem value={AVAILABLE_LANGUAGES.PORTUGUESE}>Português</SelectItem>
      <SelectItem value={AVAILABLE_LANGUAGES.ENGLISH}>English</SelectItem>
    </SelectContent>
  </Select>
  
}