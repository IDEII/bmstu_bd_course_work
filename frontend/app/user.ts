export class User {
    constructor(
        public user_profile_id: string,
        public username: string,
        public profile_type: number,
        public e_mail: string,
        public linked_id_1: string | null,
        public linked_id_2: string | null,
        public linked_id_3: string | null,
        public password: string,
    ){
    }
}