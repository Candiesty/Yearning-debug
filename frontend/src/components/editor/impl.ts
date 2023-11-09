import * as monaco from 'monaco-editor/esm/vs/editor/editor.api';
import { MysqlKeywords } from '@/components/editor/keyword';
import { lightWord } from '@/store/module/highlight';

const createSQLToken = (
  range: any,
  exact: lightWord[]
): monaco.languages.CompletionItem[] => {
  const token = [] as any;
  MysqlKeywords.forEach((item: string) => {
    token.push({
      label: `"${item}"`,
      kind: monaco.languages.CompletionItemKind.Keyword,
      insertText: `${item}`,
      range: range,
      detail: '关键字',
    });
  });

  exact.forEach((item) => {
    token.push({
      label: item.vl,
      kind: monaco.languages.CompletionItemKind.Field,
      insertText: item.vl,
      range: range,
      detail: item.meta,
    });
  });
  return token;
};

export { createSQLToken };
