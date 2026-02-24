export class Fighter {
    constructor(
        public user_id: string,
        public club_id: string | null,
        public id: string,
        public name: string,
        public description: string,
        public country: string,
        public birthday: Date,
        public rating: string,
        public category: number,
) { }
}
