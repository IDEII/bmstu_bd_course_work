import { Component, Input, OnInit } from '@angular/core';
import { TournamentService } from '../tournament.service';
import { ActivatedRoute } from '@angular/router';
import { Participant } from '../tournament';

@Component({
  selector: 'app-draw',
  templateUrl: './draw.component.html',
  styleUrl: './draw.component.css'
})
export class DrawComponent implements OnInit{
  @Input() tournamentId: string = '';
  participants: Participant[] = [];
  id : string | null = ""

  constructor(private tournamentService:  TournamentService, private route: ActivatedRoute) {
    this.participants = new Array<Participant>();
    this.id = this.route.snapshot.paramMap.get('id');
  }

  onDraw() {
    console.log(this.tournamentId)
    this.tournamentService.sendDrawRequest(this.tournamentId).subscribe(
      response => {
        console.log('Draw successful:', response);
      },
      error => {
        console.error('Error in draw:', error);
      }
    );
  }

  getParticipantName(participantId: string): string {
    const participant = this.participants.find(p => p.participantId === participantId);
    return participant ? participant.name : 'TBD';
  }

  ngOnInit(): void {
    if (this.tournamentId) {

    this.tournamentService.getParticipants(this.tournamentId).subscribe({next:(data: Participant[]) => { this.participants = data;
      console.log("yes")
    }});
    console.log(this.participants)

  }
  }
}
