import { v4 as uuidv4 } from "uuid";

export const mockUserId = uuidv4();

export const socketUrl = `ws://localhost:3000/ws/${mockUserId}`;

export const MAX_RETRY_TIMES = 5;
