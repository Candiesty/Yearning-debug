import { Timeline } from '@/apis/source';
import { request, COMMON_URI } from '@/config/request';
import { reactive, ref } from 'vue';

export default function () {
  const orderProfileArch = reactive({
    db: [] as string[],
    table: [] as string[],
    timeline: [] as Timeline[],
  });

  const editor = ref();

  const fetchStepUsage = (work_id: string) => {
    return request({
      method: 'GET',
      url: `${COMMON_URI}/fetch/steps`,
      params: {
        work_id: work_id,
      },
    });
  };

  const fetchProfileSQL = (work_id: string) => {
    return request({
      method: 'GET',
      url: `${COMMON_URI}/fetch/sql`,
      params: {
        work_id: work_id,
      },
    });
  };

  return {
    orderProfileArch,
    editor,
    fetchStepUsage,
    fetchProfileSQL,
  };
}
