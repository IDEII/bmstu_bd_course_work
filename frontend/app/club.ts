import { Fighter } from "./fighter";

export class Club {
     constructor(
       public user_id: string,
       public id: string,
       public title: string,
       public description: string | null,
       public address: string | null,
       public contact: string | null,
       public rating: string | null,
       public founded_date: Date
     ) {}
   }