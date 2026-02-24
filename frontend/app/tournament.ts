export class Tournament {
    constructor(
        public id: string,
        public name: string,
        public address: string,
        public startDate: Date, 
        public endDate: Date,
        public organazer: string,
        public rounds: string, 
        public category: number,

    ) { }
}
export class Match {
    constructor(
        public match_id: string,
        public tournament_id: string,
        public round_number: number,
        public match_number: number,
    public top_participant_id: string,
    public top_score: number,
    public bottom_participant_id: string,
    public bottom_score: number,
    public winner_id: string,
    public when_played: Date,
) { }
  }
  
export class Round {
    constructor(
        public round_number: number,
    public matches: Match[],){ }
}
  
export class TournamentBracket {
    constructor(
        public id: string,
        public rounds: Round[],
    ){ }
}

export class Participant {
    constructor(
      public participantId: string,
      public name: string
    ) {}
  }
  

  export class TableResults {
    constructor(
        public tournament_id: string, 
        public roundNumber: number, 
        public match_id: string, 
        public participant_id: string, 
        public place: string
    ) { 
    }
}