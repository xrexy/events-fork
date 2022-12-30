import type { BaseRecord } from "@/pocketbase";
import { useRecord } from "./use-record";

export interface EventDto extends BaseRecord {
  title: string;
  thumbnail: string;
  description: string;
  createdBy: string;
}

export type CreateEventDto = Pick<
  EventDto,
  "title" | "thumbnail" | "description"
> &
  Partial<Pick<EventDto, "createdBy">>;

export const useEvent = () => {
  const collection = "events" as const;
  return {
    ...useRecord<EventDto, CreateEventDto>(collection),
    createThumbnailURI: (recordId: string, fileName: string) =>
      import.meta.env.VITE_API_HOST +
      `/api/files/${collection}/${recordId}/${fileName}`,
  };
};
