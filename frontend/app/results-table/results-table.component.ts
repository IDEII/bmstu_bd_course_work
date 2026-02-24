// src/app/results-table/results-table.component.ts
import { Component, OnInit, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Participant, TableResults } from '../tournament';
import { TournamentService } from '../tournament.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-results-table',
  templateUrl: './results-table.component.html',
  styleUrls: ['./results-table.component.css']
})
export class ResultsTableComponent implements OnInit {
  @Input() tournamentid: string = '';
  results: TableResults[] = [];
  errorMessage: string | null = null;
  participants: Participant[] = [];
  id : string | null = ""

  constructor(private TournamentService:  TournamentService, private route: ActivatedRoute) {
    this.participants = new Array<Participant>();
    this.id = this.route.snapshot.paramMap.get('id');
    this.results = new Array<TableResults>();
  }

  getParticipantName(participantId: string): string {
    const participant = this.participants.find(p => p.participantId === participantId);
    return participant ? participant.name : 'TBD';
  }

  ngOnInit(): void {
    if (this.id) {

    this.TournamentService.getParticipants(this.id).subscribe({next:(data: Participant[]) => { this.participants = data;
      console.log("yes")
    }});
    console.log(this.participants)

    this.fetchResults();
  }

  }

  fetchResults() {


    if (this.id) {
    this.TournamentService.fetchResults(this.id)
      .subscribe({
        next: (data) => {
          this.results = data;
        },
        error: (error) => {
          this.errorMessage = error;
          console.error('Ошибка при получении данных:', error);
        }
      });}
  }
}