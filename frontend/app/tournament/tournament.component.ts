import { Component, Input ,OnInit, TemplateRef, ViewChild} from '@angular/core';
import { FormBuilder, FormGroup, Validators, AbstractControl, FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute, Router } from '@angular/router';
import { Tournament, Round, Match, TournamentBracket, Participant, TableResults } from '../tournament'
import { TournamentService } from '../tournament.service'

@Component({
  selector: 'app-tournament',
  templateUrl: './tournament.component.html',
  styleUrl: './tournament.component.css'
})

export class TournamentComponent {
  @Input() editable: boolean = false;
  @Input() tournamentID: string ="";
  @ViewChild("readOnlyTemplate", {static: false}) readOnlyTemplate: TemplateRef<any>|undefined;
  @ViewChild("editTemplate", {static: false}) editTemplate: TemplateRef<any>|undefined;
  participants: Participant[] = [];
  tournament : Tournament = {
    id: '',
    name: '',
    address: '',
    startDate: new Date(), 
    endDate: new Date(),
    organazer: '',
    rounds: '',
    category: 0,
  };

  tournamentBracket: TournamentBracket = {
    id: '',
    rounds: []
  }

  match : Match = {
    match_id: '',
    tournament_id: '',
    round_number: 0,
    match_number: 0,
    top_participant_id: '',
    top_score: 0,
    bottom_score: 0,
    bottom_participant_id: '',
    winner_id: '',
    when_played: new Date(),
  }
  isEditing: boolean = false; 
  oldScores: { [matchId: string]: { top: number, bottom: number } } = {}; 
  matches: Match[] = []; 

  constructor(
    private route: ActivatedRoute,
    private TournamentService: TournamentService
  ) {
    this.matches = new Array<Match>();
    this.participants = new Array<Participant>();

  }

  getParticipantName(participantId: string): string {
    const participant = this.participants.find(p => p.participantId === participantId);
    return participant ? participant.name : 'TBD';
  }
  isWinner(participantId: string, winnerId: string): boolean {

    return participantId === winnerId;
  }

  toggleEdit(): void {
    this.isEditing = !this.isEditing;

    if (this.isEditing) {
      this.tournamentBracket.rounds.forEach(round => {
        round.matches.forEach(match => {
          this.oldScores[match.match_id] = {
            top: match.top_score,               
            bottom: match.bottom_score
          };
        });
      });
      this.TournamentService.getMatchesById(this.tournament.id).subscribe(data => {
        this.matches = data; 
        console.log("Матчи загружены:", this.matches);
      });                                             
    } else {
      this.oldScores = {};
    }
  }

  update(match: Match): void {
        const matchID = match.match_id; 
        const topScore = match.top_score;
        const bottomScore = match.bottom_score; 

        if (this.isMatchReadyToUpdate(match)) {
          console.log(`Обновление для матча ${matchID}: ${topScore}, ${bottomScore}`);
          
          this.TournamentService.updateMatchScore(matchID, topScore, bottomScore, this.tournament.id)
            .subscribe(response => {
              console.log('Счет успешно обновлен:', response);
            }, error => {
              console.error('Ошибка при обновлении счета:', error);
            });
        } else {
          console.log(`Матч ${matchID} пропущен: неверные условия для обновления.`);
        }
      
    this.TournamentService.getBracketData(this.tournament.id).subscribe(data => {
      this.tournamentBracket = data;
    });
  }

  getScore(matchId: string, participantId: string): number {
    const match = this.matches.find(m => m.match_id === matchId);
    console.log(matchId)
    console.log(match?.match_id)
    console.log(match?.bottom_score)

    if (match) {
      if (participantId === match.top_participant_id) {
        return match.top_score;
      } else if (participantId === match.bottom_participant_id) {
        return match.bottom_score;
      }
    }
    return 0;
  }
  
  isMatchReadyToUpdate(match: any): boolean {
    return match.top_participant_id && 
           match.bottom_participant_id && 
           match.top_score !== undefined && 
           match.bottom_score !== undefined &&
           (match.top_score !== this.oldScores[match.match_id]?.top || 
            match.bottom_score !== this.oldScores[match.match_id]?.bottom);
  }

  ngOnInit() {
    var id = this.route.snapshot.paramMap.get('id');
    if (id !== this.tournamentID && this.tournamentID !== "") {
      id = this.tournamentID
      this.editable = true;
    }
    if (id !== null) {
    this.TournamentService.getTournament(id).subscribe(data => {
      this.tournament = data;
    });
    this.TournamentService.getBracketData(id).subscribe(data => {
      this.tournamentBracket = data;
    });
    this.TournamentService.getParticipants(id).subscribe({next:(data: Participant[]) => { this.participants = data; 
      console.log("yes")
    }});
    this.TournamentService.getMatchesById(id).subscribe(data => {
      this.matches = data;
      console.log("Матчи загружены:", this.matches);
    });
  }
  }
  
}
