export class Request {
    constructor(
        public request_id: string,
        public id_fighter: string,
        public tournament_id: string,
        public invited_id: string,
        public decision: number,
    ){ }
}

export class InvitedParticipant {
    constructor(
        public fighter_id: string,
        public tournament_id: string,
        public InvitedParticipant_id: string,
    ){

    }
}

export class RequestWithTournamentTitle {
    constructor(
        public request_id: string,
        public id_fighter: string,
        public name: string,
        public tournament_id: string,
        public category: number,
        public decision: number 
    ) { }
}
